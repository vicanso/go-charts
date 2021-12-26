var height = document.body.clientHeight- 110;
var editor = CodeMirror.fromTextArea(document.getElementById("codeInput"), {
    lineNumbers: true,
    lineWrapping: true,
    mode: "javascript"
});
editor.setSize("100%", height);

function run() {
    var option = editor.getValue();
    var data = null;
    try {
        if (option.indexOf("option = ") !== -1) {
            var fn = new Function("var " + option + ";return option;");
            data = fn();
        } else {
            data = JSON.parse(option);
        }
    } catch (err) {
        alert(err.message);
        return;
    }

    axios.post("/", data).then(function(resp) {
        document.getElementById("svg").innerHTML = resp;
    }).catch(function(err) {
        alert(err.message);
    });
}