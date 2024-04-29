from flask import Flask, request
import torch
from model import MyModel

app = Flask(__name__)
model_path = "my_model.pth"

# Load the model
model = MyModel()  # replace this with your model class
model.load_state_dict(torch.load(model_path))
model.eval()


@app.route("/predict", methods=["POST"])
def predict():
    data = request.get_json()  # get the request data
    input_tensor = torch.tensor(data["input"])  # convert it to a tensor
    output = model(input_tensor)  # run the model
    return {"output": output.tolist()}  # return the output


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
