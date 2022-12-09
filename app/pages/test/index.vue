<template>
    <p>test</p>
</template>
<script lang="ts" setup>
onMounted( async() => {
    const config = useRuntimeConfig()
    const ws:WebSocket = new WebSocket(config.socket)
    ws.onopen = function () {
        console.log("接続が開かれたときに呼び出されるイベント");
    };
    ws.onmessage = function (event) {
        console.log(event.data)
    };
    function send() {
        let send_msg = "aaaa"
        ws.send(send_msg);
    }
})
</script>
