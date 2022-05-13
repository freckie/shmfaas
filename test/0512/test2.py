import os, psutil

proc = psutil.Process(os.getpid())
curr = 0
msgfmt = '$ %10d | %s'

def mem(msg=''):
    print(msgfmt % (proc.memory_info().rss, msg))

import torch
import torchvision.models as models
from torchvision.transforms import ToTensor

import numpy as np
from PIL import Image

import pickle
from multiprocessing.shared_memory import SharedMemory

import xtorch

if __name__ == '__main__':
    mem('Initilized')

    # model load
    model_skeleton = models.vgg16(False, False)
    mem('After loading model')

    # load metadata from pickle
    metadata: xtorch.XMetadata
    with open('/Users/freckie/prj/shmfaas/torchtest/0512-test/vgg16-meta', 'rb') as f:
        metadata = pickle.load(f)
    mem('After loading the metadata')

    # load tensors into the model
    shm, model = xtorch.x_load_states(model_skeleton, metadata)
    model.eval()
    mem('After loading tensors')

    # del the skeleton model
    del model_skeleton
    del metadata
    mem('After releasing the skeleton model')

    input()

    # predict
    input_img = Image.open('/Users/freckie/prj/shmfaas/torchtest/dog-224.jpg').convert('RGB')
    input_tensor = torch.unsqueeze(ToTensor()(np.array(input_img)), 0)
    model.eval()
    result = model(input_tensor)
    print(torch.argmax(result, dim=1))

    input()

    shm.close()
        