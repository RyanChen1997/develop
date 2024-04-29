"""安装推理所需要的包、下载Demo，使用方法：python init_infer_env.py --help
usage: init_infer_env.py [-h] [--tiinfer-version {2.0}] --model-format {TorchScript, Detectron2, MMDetection, ONNX, Tensorflow, PMML} --framework {tf1.15.0, tf2.4.0, torch1.9.0, torch1.12.1, onnx1.11.0, jpmml0.6.2}
                         [--model-scene {detect, classify}][--demo-dir DEMO_DIR] [--index-url INDEX_URL]

安装推理所需要的包、下载Demo。 eg: python3 init_infer_env.py --framework torch1.9.0 --model-format Detectron2 --model-scene detect --demo-dir ./
平台支持的模型包格式：
TorchScript : torch1.9.0/torch1.12.1
Detectron2 : torch1.9.0
ONNX : onnx1.11.0
MMDetection : torch1.9.0
Tensorflow: tf1.15.0/tf2.4.0
PMML: jpmml0.6.2

optional arguments:
  -h, --help            show this help message and exit
  --tiinfer-version {2.0}, -v {2.0} ，推理服务框架版本，非必填，默认最新，支持: ['2.0']
  --model-format {TorchScript,Detectron2,ONNX,MMDetection,Tensorflow,PMML}, -m {TorchScript,Detectron2,ONNX,MMDetection,Tensorflow,PMML} 模型格式，必填，平台目前支持以下几种：['TorchScript', 'Detectron2', 'ONNX', 'MMDetection', 'Tensorflow', 'PMML']，请输入其中的一种
  --framework {tf1.15.0,tf2.4.0,torch1.9.0,torch1.12.1}, -f {tf1.15.0,tf2.4.0,torch1.9.0,torch1.12.1} 推理框架，必填，支持: ['tf1.15.0', 'tf2.4.0', 'torch1.9.0', 'torch1.12.1']
  --model-scene {detect,classify} -s {detect, classify} 模型使用场景，非必填，默认'detect'，支持：['detect', 'classify']
  --demo-dir DEMO_DIR, -d DEMO_DIR 下载的路径，非必填，默认不做下载
  --index-url INDEX_URL, -i INDEX_URL 镜像源地址
"""
import os
import textwrap

from pip._internal import main
import argparse
import logging
import textwrap as _textwrap

logging.basicConfig(
    level=logging.INFO, format="%(asctime)s [%(levelname)s] %(message)s"
)

# 支持的版本以及框架对应的python包和demo地址
Config = {
    "ti_infer_versions": ["1.0"],
    "model_format": [
        "TorchScript",
        "Detectron2",
        "ONNX",
        "FrozenGraph",
        "SavedModel",
        "MMDetection",
        "PMML",
        "HuggingFace-Bert",
        "HuggingFace-StableDiffusion",
        "HuggingFace-StableDiffusion-DynamicLora",
    ],
    "frameworks": {
        "torch1.9.0": {
            "TorchScript": {
                "packages": [
                    "torch==1.9.0",
                    "opencv-contrib-python==4.6.0.66",
                    "opencv-python==4.6.0.66",
                    "setuptools==59.5.0",
                    "mmcv-full==1.4.8",
                    "transformers==4.19.4",
                    "torchvision==0.10.0",
                ],
                "cmds": ["pip3 install easyocr==1.6.2"],
                "demo_url": {
                    "detect": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/torch_script/detect.zip",
                    "classify": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/torch_script/classify.zip",
                    "nlp": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/torch_script/nlp.zip",
                    "ocr": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/torch_script/ocr.zip",
                },
            },
            "HuggingFace-Bert": {
                "packages": [
                    "torch==1.9.0",
                    "opencv-contrib-python==4.6.0.66",
                    "opencv-python==4.6.0.66",
                    "setuptools==59.5.0",
                    "mmcv-full==1.4.8",
                    "transformers==4.19.4",
                    "torchvision==0.10.0",
                ],
                "cmds": [],
                "demo_url": {
                    "nlp": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/hugging_face/nlp.zip"
                },
            },
            "Detectron2": {
                "packages": ["opencv-contrib-python==4.6.0.66", "setuptools==59.5.0"],
                "cmds": [
                    "pip install --no-cache-dir torch==1.9.0+cu111 torchvision==0.10.0+cu111 --extra-index-url https://download.pytorch.org/whl/cu111",
                    "python -m pip install detectron2 -f https://dl.fbaipublicfiles.com/detectron2/wheels/cu111/torch1.9/index.html",
                ],
                "demo_url": {
                    "detect": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/detectron2/detect.zip"
                },
            },
            "MMDetection": {
                "packages": [
                    "opencv-contrib-python==4.6.0.66",
                    "opencv-python==4.6.0.66",
                    "setuptools==59.5.0",
                    "transformers==4.19.4",
                    "pycocotools",
                ],
                "cmds": [
                    "pip3 install --no-cache-dir torch==1.9.0+cu111 torchvision==0.10.0+cu111 torchaudio==0.9.0 -f https://download.pytorch.org/whl/torch_stable.html",
                    "pip3 install --no-cache-dir mmcv-full==1.4.8 -f https://download.openmmlab.com/mmcv/dist/cu111/torch1.9.0/index.html",
                ],
                "demo_url": {
                    "detect": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/mmdetection/detect.zip"
                },
            },
        },
        "torch1.12.1": {
            "TorchScript": {
                "packages": [
                    "torch==1.12.1",
                    "opencv-contrib-python==4.6.0.66",
                    "opencv-python==4.6.0.66",
                    "mmcv-full==1.4.8",
                    "transformers==4.23.0",
                    "torchvision==0.13.1",
                ],
                "cmds": ["pip3 install easyocr==1.6.2"],
                "demo_url": {
                    "detect": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/torch_script/detect.zip",
                    "classify": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/torch_script/classify.zip",
                    "nlp": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/torch_script/nlp.zip",
                    "ocr": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/torch_script/ocr.zip",
                },
            },
        },
        "torch2.0.0": {
            "HuggingFace-StableDiffusion": {
                "packages": [
                    "torch==2.0.0",
                    "torchvision==0.15.1",
                    "opencv-contrib-python==4.6.0.66",
                    "opencv-python==4.6.0.66",
                    "setuptools==59.5.0",
                    "transformers==4.26.1",
                    "diffusers==0.16.1",
                    "accelerate==0.16.0",
                    "safetensors==0.3.1",
                    "protobuf==3.20.3",
                ],
                "cmds": [],
                "demo_url": {
                    "sd": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/hugging_face/sd.zip",
                    "sd_cn": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/hugging_face/sd_cn.zip",
                    "sd_lora": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/hugging_face/sd_lora.zip",
                },
            },
            "HuggingFace-StableDiffusion-DynamicLora": {
                "packages": [
                    "torch==2.0.0",
                    "torchvision==0.15.1",
                    "opencv-contrib-python==4.6.0.66",
                    "opencv-python==4.6.0.66",
                    "setuptools==59.5.0",
                    "transformers==4.26.1",
                    "diffusers==0.16.1",
                    "accelerate==0.16.0",
                    "safetensors==0.3.1",
                    "protobuf==3.20.1",
                ],
                "cmds": [],
                "demo_url": {
                    "sd_lora": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/hugging_face/sd_lora.zip",
                },
            },
        },
        "onnx1.11.0": {
            "ONNX": {
                "packages": [
                    "torch==1.9.0",
                    "onnx==1.11.0",
                    "onnxruntime-gpu==1.11.1",
                    "opencv-contrib-python==4.6.0.66",
                    "setuptools==59.5.0",
                    "torchvision==0.10.0",
                    "scipy==1.10.1",
                    "shapely==2.0.1",
                    "pyclipper",
                    "python-box",
                    "six",
                ],
                "cmds": [],
                "demo_url": {
                    "detect": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/onnx/detect.zip"
                },
            }
        },
        "tf1.15.0": {
            "FrozenGraph": {
                "packages": [
                    "tensorflow-gpu==1.15.0",
                    "opencv-contrib-python==4.6.0.66",
                    "setuptools==59.5.0",
                    "protobuf==3.20.1",
                ],
                "cmds": [],
                "demo_url": {
                    "nlp": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/frozen_graph/nlp.zip"
                },
            }
        },
        "tf2.4.0": {
            "FrozenGraph": {
                "packages": [
                    "tensorflow-gpu==2.4.0",
                    "opencv-contrib-python==4.6.0.66",
                    "setuptools==59.5.0",
                ],
                "cmds": [],
                "demo_url": {
                    "nlp": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/frozen_graph/nlp.zip"
                },
            },
            "SavedModel": {
                "packages": [
                    "tensorflow-gpu==2.4.0",
                    "opencv-contrib-python==4.6.0.66",
                    "setuptools==59.5.0",
                ],
                "cmds": [],
                "demo_url": {
                    "nlp": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/saved_model/nlp.zip",
                    "recommend": "https://tione-dev-1256580188.cos.ap-guangzhou.myqcloud.com/ti-cloud/ti-ems/ti-infer-demo/saved_model/recommend.zip",
                },
            },
        },
        "jpmml0.6.2": {
            "PMML": {"packages": ["jpmml_evaluator==0.6.2"], "cmds": [], "demo_url": {}}
        },
    },
    "framework_versions": [
        "torch1.9.0",
        "torch1.12.1",
        "tf1.15.0",
        "tf2.4.0",
        "jpmml0.6.2",
        "onnx1.11.0",
    ],
    "model_scene": ["detect", "classify", "nlp", "ocr", "recommend", "sd"],
}


# 默认的镜像源地址
# DefaultIndex = "http://mirrors.tencent.com/pypi/simple"
DefaultIndex = None
# 默认模型场景
DefaultScene = "detect"


def get_host(index_url):
    if index_url is None:
        index_url = index_url.replace("https://", "").replace("http://", "")
        index = index_url.index("/")
        if index > 0:
            return index_url[:index]

    return index_url


def install_ti_infer(version, index_url):
    """
    安装tiinfer推理服务相关依赖包
    :param version: tiinfer的版本号
    :type version:  str
    :param index_url: 镜像源地址
    :type index_url:  str
    :return: 空
    :rtype:
    """
    ti_infer_package = "ti-cloud-infer-framework"
    mosce_ti_infer_package = "mosec-tiinfer"
    if version is not None:
        ti_infer_package += "==" + version
        mosce_ti_infer_package += "==" + version

    if index_url is None:
        install_args = [
            "install",
            ti_infer_package,
            mosce_ti_infer_package,
            "requests>=2.26.0",
            "urllib3>=1.26.7",
            "PyYAML>=5.4.1",
            "certifi>=2021.10.8",
            "easydict",
            "tqdm",
            "terminaltables",
            "tabulate",
        ]
    else:
        install_args = [
            "install",
            "-i",
            index_url,
            "--trusted-host",
            get_host(index_url),
            ti_infer_package,
            mosce_ti_infer_package,
            "requests>=2.26.0",
            "urllib3>=1.26.7",
            "PyYAML>=5.4.1",
            "certifi>=2021.10.8",
            "easydict",
            "tqdm",
            "terminaltables",
            "tabulate",
        ]

    logging.info(install_args)
    ret = main(install_args)
    if ret != 0:
        raise Exception("install_ti_infer failed")


def download_unzip_file(url, save_path):
    """
    下载并解压文件
    :param url: 下载的文件的url
    :type url:
    :param save_path: 保存的地址
    :type save_path:  str
    :return: 空
    :rtype:
    """
    logging.info("download url: %s, save dir: %s", url, save_path)
    import requests
    import zipfile
    import tempfile

    rsp = requests.get(url)

    _tmp_file = tempfile.TemporaryFile()
    _tmp_file.write(rsp.content)
    zf = zipfile.ZipFile(_tmp_file, mode="r")
    for names in zf.namelist():
        f = zf.extract(names, save_path)  # 解压到zip目录文件下
        logging.info(f)
    zf.close()


def install_framework(framework_name, model_format, demo_dir, index_url, model_scene):
    """
    安装推理推理框架（tensorflow，pytorch等）相关依赖包，以及下载对应的demo
    :param framework_name: 框架的名称
    :type framework_name:  str
    :param demo_dir: 下载的demo的路径
    :type demo_dir:  str
    :param index_url: 镜像源地址
    :type index_url:  str
    :parm model_format: 模型格式
    :type model_format: str
    :parm model_scene: 模型使用场景
    :type model_scene: str
    :return: 空
    :rtype:
    """
    if index_url is None:
        install_args = ["install"]
    else:
        install_args = [
            "install",
            "-i",
            index_url,
            "--trusted-host",
            get_host(index_url),
        ]

    packages = Config["frameworks"][framework_name][model_format]["packages"]
    install_args += packages

    logging.info(install_args)
    ret = main(install_args)

    cmds = Config["frameworks"][framework_name][model_format]["cmds"]
    for i in cmds:
        logging.info(i)
        f = os.popen(i)
        logging.info(f.read())

    if ret != 0:
        raise Exception("install_framework failed")
    if demo_dir is not None:
        download_unzip_file(
            Config["frameworks"][framework_name][model_format]["demo_url"][model_scene],
            demo_dir,
        )


class MultilineFormatter(argparse.HelpFormatter):
    def _fill_text(self, text, width, indent):
        text = self._whitespace_matcher.sub(" ", text).strip()
        paragraphs = text.split("|n ")
        multiline_text = ""
        for paragraph in paragraphs:
            formatted_paragraph = (
                _textwrap.fill(
                    paragraph, width, initial_indent=indent, subsequent_indent=indent
                )
                + "\n"
            )
            multiline_text = multiline_text + formatted_paragraph
        return multiline_text


if __name__ == "__main__":

    # 方法参数说明和help说明
    parser = argparse.ArgumentParser(
        description="""安装推理所需要的依赖包、下载Demo。 |n
                                     eg: python init_infer_env.py --framework torch1.9.0 --model-format TorchScript --model-scene detect --demo-dir ./ |n
                                     平台支持的模型包格式： |n
                                     TorchScript  : ti-cloud-infer-pytorch-gpu:py38-torch1.9.0-cu111-3.0.3/ti-cloud-infer-pytorch-gpu:py38-torch1.12.1-cu113-3.0.4 |n
                                     Detectron2   : ti-cloud-infer-pytorch-gpu:py38-torch1.9.0-detectron2-cu111-3.0.3 |n
                                     ONNX         : ti-cloud-infer-pytorch-gpu:py38-torch1.9.0-onnx1.11.1-cu111-3.0.1 |n
                                     MMDetection  : ti-cloud-infer-pytorch-gpu:py38-torch1.9.0-mmdetection-cu111-3.0.3 |n
                                     FrozenGraph  : ti-cloud-infer-tensorflow-gpu:py38-tf2.4-cu11.0-2.0./ti-cloud-infer-tensorflow-gpu:py37-tf1.15-cu10.0-2.0.1 |n
                                     Savedmodel   : ti-cloud-infer-tensorflow-gpu:py38-tf2.4-cu11.0-2.0. |n
                                     PyTorch      : ti-cloud-infer-pytorch-gpu:py38-torch1.9.0-cu111-3.0.3/ti-cloud-infer-pytorch-gpu:py38-torch1.12.1-cu113-3.0.4 |n
                                     HuggingFace-Bert  : ti-cloud-infer-pytorch-gpu:py38-torch1.9.0-cu111-3.0.3 |n
                                     PMML         : ti-cloud-infer-pmml:py38-jpmml0.6.2-2.0.1 |n
                                     """,
        formatter_class=MultilineFormatter,
    )
    parser.add_argument(
        "--ti-infer-version",
        "-v",
        help="推理服务框架版本，非必填，默认最新，支持: " + Config["ti_infer_versions"].__str__(),
        choices=Config["ti_infer_versions"],
    )
    parser.add_argument(
        "--model-format",
        "-m",
        help="模型格式，必填，平台目前支持以下几种：" + Config["model_format"].__str__() + "，请输入其中的一种",
        choices=Config["model_format"],
        required=True,
    )
    parser.add_argument(
        "--framework",
        "-f",
        help="推理框架，必填，torch或者tensorflow的版本，必须与模型格式对应，支持:"
        + Config["framework_versions"].__str__(),
        choices=list(Config["frameworks"].keys()),
        required=True,
    )
    parser.add_argument("--demo-dir", "-d", help="下载的路径，非必填，默认不做下载")
    parser.add_argument("--index-url", "-i", help="镜像源地址，非必填", default=DefaultIndex)
    parser.add_argument(
        "--model-scene",
        "-s",
        help="模型使用场景，非必填，默认detect，支持：" + Config["model_scene"].__str__(),
        choices=Config["model_scene"],
        default=DefaultScene,
    )
    args = parser.parse_args()

    # 安装tiinfer推理服务相关依赖包
    install_ti_infer(args.ti_infer_version, DefaultIndex)

    # 安装推理推理框架（tensorflow，pytorch等）相关依赖包，以及下载对应的demo
    install_framework(
        args.framework,
        args.model_format,
        args.demo_dir,
        args.index_url,
        args.model_scene,
    )
