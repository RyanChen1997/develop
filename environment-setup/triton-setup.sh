#!/bin/bash

current_dir=`pwd`
triton_dir=$current_dir"/triton"
model_dir=$triton_dir"/model_repository"

# create a model repository
mkdir triton
cd $triton_dir
wget https://raw.githubusercontent.com/triton-inference-server/server/main/docs/examples/fetch_models.sh
bash fetch_models.sh


# launch triton
docker run -d --rm -p8000:8000 -p8001:8001 -p8002:8002 \
	-v$model_dir:/models \
	nvcr.io/nvidia/tritonserver:22.07-py3 tritonserver \
	--model-repository=/models --model-control-mode=poll

# test model service
docker run --rm --net=host nvcr.io/nvidia/tritonserver:22.07-py3-sdk /workspace/install/bin/image_client -m densenet_onnx -c 3 -s INCEPTION /workspace/images/mug.jpg
