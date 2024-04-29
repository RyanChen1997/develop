import torch
import torch.nn as nn
import torch.optim as optim

# Define the model
class MyModel(nn.Module):
    def __init__(self):
        super(MyModel, self).__init__()
        self.linear = nn.Linear(1, 1)  # One in and one out

    def forward(self, x):
        return self.linear(x)

# Create the model
model = MyModel()

# Define loss function and optimizer
criterion = nn.MSELoss()
optimizer = optim.SGD(model.parameters(), lr=0.01)

# Training data (x = y - 1)
x_train = torch.tensor([[0.], [1.], [2.], [3.], [4.]], requires_grad=True)
y_train = torch.tensor([[1.], [2.], [3.], [4.], [5.]], requires_grad=True)

# Train the model
for epoch in range(1000):
    model.zero_grad()
    outputs = model(x_train)
    loss = criterion(outputs, y_train)
    loss.backward()
    optimizer.step()

# Save the trained model
torch.save(model.state_dict(), 'my_model.pth')

# Test the model
x_test = torch.tensor([[5.]], requires_grad=True)
print(model(x_test))  # Should output a value close to 6
