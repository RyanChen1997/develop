// theater
const inputUrl = document.querySelector(".theater input[id='video-url']");
const button = document.querySelector(".theater button");
const theater = document.querySelector(".theater");

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

function extractM3U8FromUrl(url) {
    if (url === '') {
        return;
    }

    const requestURL = "http://localhost:8080/extract?url=" + url;
    fetch(requestURL).
        then(response => response.json()).
        then(data => {
            if (!data?.urls) return;
            addUrls(theater, data.urls);
        }).
        catch(error => console.log(error))

}

function addUrls(theater, urls) {
    if (urls.length === 0) {
        return;
    }

    var ul = document.createElement("ul");
    urls.forEach(url => {
        const li = document.createElement("li");
        li.textContent = url;
        ul.appendChild(li);
    });

    theater.appendChild(ul);
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
