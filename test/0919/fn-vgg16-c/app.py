from multiprocessing.shared_memory import SharedMemory
import os
import json
import pickle

from flask import Flask

import numpy as np
import torch
import torchvision.models as models
from torchvision.transforms import ToTensor
from PIL import Image

import shmtorch

shm: SharedMemory
app = Flask(__name__)

@app.route('/')
def predict():
    env = os.environ
    addr = env['NODE_NAME'] + ':20000'
    model_name = env['SHMM_NAME']
    tag_name = env['TAG_NAME']

    model_skeleton = models.vgg16(False, False)

    metadata: shmtorch.XMetadata
    metadata = shmtorch.x_get_metadata(addr, model_name, tag_name)
    
    shm, model = shmtorch.x_load_states(model_skeleton, metadata)
    model.eval()

    del model_skeleton
    del metadata

    img = Image.open('/app/dog-224.jpg').convert('RGB')
    t = torch.unsqueeze(ToTensor()(np.array(img)), 0)
    result = model(t)
    result_json = json.dumps({'result': str(torch.argmax(result, dim=1))})
    shm.close()

    return result_json

if __name__ == "__main__":
    try:
        app.run(debug=True,host='0.0.0.0',port=int(os.environ.get('PORT', 8080)))
    except Exception as e:
        print('Exception on __main__ :', e)
        shm.close()
