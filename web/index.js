var height = document.body.clientHeight- 110;
var editor = CodeMirror.fromTextArea(document.getElementById("codeInput"), {
    lineNumbers: true,
    lineWrapping: true,
    mode: "javascript"
});
editor.setSize("100%", height);

function run() {
    var option = editor.getValue();
    axios.post("/", JSON.parse(option)).then(function(resp) {
        document.getElementById("svg").innerHTML = resp;
    });
}