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

@app.route('/vgg16')
def predict():
    model_skeleton = models.vgg16(False, False)

    metadata: shmtorch.XMetadata
    with open('/app/metadata', 'rb') as f:
        metadata = pickle.load(f)

    shm, model = shmtorch.x_load_states(model_skeleton, metadata)
    model.eval()

    del model_skeleton
    del metadata

    img = Image.open('/app/dog-224.jpg').convert('RGB')
    t = torch.unsqueeze(ToTensor()(np.array(img)), 0)
    result = model(t)

    return json.dumps({'result': str(torch.argmax(result, dim=1))})

if __name__ == "__main__":
    app.run(debug=True,host='0.0.0.0',port=int(os.environ.get('PORT', 8080)))