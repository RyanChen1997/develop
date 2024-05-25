const payElement = document.querySelector(".header #pay")
const coffeeElement = document.querySelector(".header .buy-me-coffee")

// 为元素及其所有子元素添加事件监听器
function addListenerToElementAndChildren(eventType, element, eventHandler) {
    element.addEventListener(eventType, eventHandler);

    var children = element.children;
    for (var i = 0; i < children.length; i++) {
        addListenerToElementAndChildren(eventType, children[i], eventHandler);
    }
}

document.addEventListener("click", function (e) {
    if (e.target.id == "pay" || e.target.parentNode.id == "pay"
        || e.target.className == "buy-me-coffee" || e.target.parentNode.className == "buy-me-coffee") {
        return;
    }
    payElement.style.display = "none";
});

addListenerToElementAndChildren("click", coffeeElement, function (e) {
    payElement.style.display = "inline-block";
});