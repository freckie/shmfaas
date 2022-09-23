import os

import torch
import numpy as np
import torchvision.models as models

if __name__ == "__main__":
    model = models.vgg16(True, True)
    torch.save(model.state_dict(), './vgg16.pth')

    print('Finished')
