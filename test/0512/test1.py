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

import shmtorch

if __name__ == '__main__':
    mem('Initilized')

    # model load
    model = models.vgg16(False, False)
    mem('After loading model')

    # load state_dict
    state_dict = torch.load('../vgg16.pth')
    mem('After loading state_dict from pth file')

    model.load_state_dict(state_dict)
    mem('After connecting state_dict and model')

    # save to shm
    shm, metadata = shmtorch.x_save_states(model, 'shm_1705')
    with open('./vgg16-meta', 'wb') as f:
        pickle.dump(metadata, f)
    print(metadata)
    mem('After saving tensors to shm')

    # del state_dict
    del state_dict
    mem('Deleted state_dict')

    input()

    # shm unlink
    shm.close()
    shm.unlink()