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
    model = models.vgg16(False, False)
    mem('After loading model')

    # load state_dict
    state_dict = torch.load('/Users/freckie/prj/shmfaas/torchtest/vgg16.pth')
    mem('After loading state_dict')

    # set state_dict to model
    model.load_state_dict(state_dict)
    mem('After connecting model and state_dict')

    del state_dict
    mem('After releasing state_dict')

    input()

    # predict
    input_img = Image.open('../dog-224.jpg').convert('RGB')
    input_tensor = torch.unsqueeze(ToTensor()(np.array(input_img)), 0)
    model.eval()
    result = model(input_tensor)
    print(torch.argmax(result, dim=1))