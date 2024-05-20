// theater
const inputUrl = document.querySelector(".theater input[id='video-url']");
const button = document.querySelector(".theater button");
const theater = document.querySelector(".theater");
const loading = document.querySelector(".theater .loading");
const result = document.querySelector(".theater .result");
const warnElement = document.querySelector(".theater .warn");
const theaterSearchElement = theater.querySelector(".theater-search");
const theaterPlayElement = theater.querySelector(".theater-play");

const remoteHost = 'http://101.126.9.35'

inputUrl.addEventListener('keydown', function (e) {
    if (e.key === 'Enter') {
        getAndClearInput(inputUrl, extractM3U8FromUrl)
    }
})

button.addEventListener("click", function (e) {
    getAndClearInput(inputUrl, extractM3U8FromUrl)
})

function getAndClearInput(input, callback) {
    const value = input.value;
    input.value = '';
    callback(value);
}

function userCanInput(enable) {
    if (enable) {
        button.disabled = false;
        inputUrl.disabled = false;
    } else {
        button.disabled = true;
        inputUrl.disabled = true;
    }
}

function extractM3U8FromUrl(url) {
    if (url === '') {
        return;
    }

    theaterMainContentShow("loading");

    const requestURL = remoteHost + "/extract?url=" + url;
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
            setUrls(theaterPlayElement, data.urls);
            theaterMainContentShow("result")
        }).
        catch(error => {
            console.log(error);
            warn(error);
            theaterMainContentShow("warn");
        });
}

function theaterMainContentShow(elem) {
    theaterPlayElement.innerHTML = '';

    if (elem === "none") {
        loading.style.display = "none";
        result.style.display = "none";
        warnElement.style.display = "none";
        userCanInput(true);
    }
    if (elem === "loading") {
        loading.style.display = "block";
        result.style.display = "none";
        warnElement.style.display = "none";
        userCanInput(false);
    }
    if (elem === "result") {
        result.style.display = "block";
        warnElement.style.display = "none";
        loading.style.display = "none";
        userCanInput(true);
    }
    if (elem === "warn") {
        warnElement.style.display = "block";
        loading.style.display = "none";
        result.style.display = "none";
        userCanInput(true);
    }
}

function warn(text) {
    warnElement.querySelector("p").innerText = text;
}

function clickUrlPlayVideo(li) {
    li.addEventListener("click", function (e) {
        theaterPlayElement.innerHTML = '';

        const url = li.innerText.trim();
        const video = document.createElement("video");
        video.setAttribute("controls", "controls");
        theaterPlayElement.appendChild(video);
        playWithHls(video, url);
    })
}

function setUrls(theater, urls) {
    if (urls.length === 0) {
        return;
    }

    var ul = result.querySelector("ul");
    ul.innerHTML = '';

    urls.forEach(url => {
        const li = document.createElement("li");
        li.innerHTML = '<i class="fa-solid fa-hand-point-right"></i> {url}'.replace("{url}", url);
        ul.appendChild(li);
    });

    ul.querySelectorAll("li").forEach(li => clickUrlPlayVideo(li))
}


function playWithHls(video, url) {
    var videoSrc = url;
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
