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

app = Flask(__name__)

@app.route('/')
def predict():
    model = models.vgg16(False, False)
    model.load_state_dict(torch.load('./vgg16.pth'))
    model.eval()

    img = Image.open('/app/dog-224.jpg').convert('RGB')
    t = torch.unsqueeze(ToTensor()(np.array(img)), 0)
    result = model(t)
    result_json = json.dumps({'result': str(torch.argmax(result, dim=1))})

    return result_json

if __name__ == "__main__":
    try:
        app.run(debug=True,host='0.0.0.0',port=int(os.environ.get('PORT', 8080)))
    except Exception as e:
        print('Exception on __main__ :', e)
