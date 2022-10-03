import os

import torch
import numpy as np
import torchvision.models as models

if __name__ == "__main__":
    model = models.resnet50(True, True)
    torch.save(model.state_dict(), './resnet50.pth')

    print('Finished')
