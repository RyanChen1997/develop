// theater
const inputUrl = document.querySelector(".theater input[id='video-url']");
const button = document.querySelector(".theater button");
const theater = document.querySelector(".theater");
const loading = document.querySelector(".theater .loading");
const result = document.querySelector(".theater .result");
const warnElement = document.querySelector(".theater .warn");

inputUrl.addEventListener('keydown', function (e) {
    if (e.key === 'Enter') {
        getAndClearInput(inputUrl, extractM3U8FromUrl)
        userCanInput(false)
    }
})

button.addEventListener("click", function (e) {
    getAndClearInput(inputUrl, extractM3U8FromUrl)
    userCanInput(false)
})

function getAndClearInput(input, callback) {
    const value = input.value;
    input.value = '';
    callback(value);
}

function userCanInput(enable) {
    if (enable) {
        button.disabled = true;
        inputUrl.disabled = true;
    } else {
        button.disabled = false;
        inputUrl.disabled = false;
    }
}

function extractM3U8FromUrl(url) {
    if (url === '') {
        return;
    }

    theaterMainContentShow("loading");

    const requestURL = "http://localhost:8080/extract?url=" + url;
    fetch(requestURL).
        then(response => {
            if (response.status !== 200) {
                console.log(response);
                warn("请求错误")
                theaterMainContentShow("warn");
                return;
            }
            return response.json();
        }).
        then(data => {
            if (!data?.urls) {
                console.log(data);
                warn("数据格式错误!")
                theaterMainContentShow("warn");
                return;
            };
            setUrls(theater, data.urls);
            theaterMainContentShow("result")
        }).
        catch(error => {
            console.log(error);
            warn(error);
            theaterMainContentShow("warn");
        }).
        finally(() => {
            userCanInput(true);
        });
}

function theaterMainContentShow(elem) {
    if (elem === "none") {
        loading.style.display = "none";
        result.style.display = "none";
        warnElement.style.display = "none";
    }
    if (elem === "loading") {
        loading.style.display = "block";
        result.style.display = "none";
        warnElement.style.display = "none";
    }
    if (elem === "result") {
        result.style.display = "block";
        warnElement.style.display = "none";
        loading.style.display = "none";
    }
    if (elem === "warn") {
        warnElement.style.display = "block";
        loading.style.display = "none";
        result.style.display = "none";
    }
}

function warn(text) {
    warnElement.querySelector("p").innerText = text;
}

function setUrls(theater, urls) {
    if (urls.length === 0) {
        return;
    }

    var ul = result.querySelector("ul");
    ul.innerHTML = '';

    //
    urls.forEach(url => {
        const li = document.createElement("li");
        li.innerHTML = '<li><i class="fa-solid fa-hand-point-right"></i> {url}</li>'.replace("{url}", url);
        ul.appendChild(li);
    });
}


function playWithHls() {
    var video = document.getElementById('video');
    var videoSrc = 'https://v8.dious.cc/20230504/5lKH0ozs/index.m3u8';
    if (Hls.isSupported()) {
        console.log('HLS is supported');
        var hls = new Hls();
        hls.loadSource(videoSrc);
        hls.attachMedia(video);
    }
    // HLS.js is not supported on platforms that do not have Media Source
    // Extensions (MSE) enabled.
    //
    // When the browser has built-in HLS support (check using `canPlayType`),
    // we can provide an HLS manifest (i.e. .m3u8 URL) directly to the video
    // element through the `src` property. This is using the built-in support
    // of the plain video element, without using HLS.js.
    else if (video.canPlayType('application/vnd.apple.mpegurl')) {
        console.log('This browser can play HLS video!');
        video.src = videoSrc;
    }
}
