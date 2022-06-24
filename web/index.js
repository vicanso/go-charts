var height = document.body.clientHeight- 110;
var editor = CodeMirror.fromTextArea(document.getElementById("codeInput"), {
    lineNumbers: true,
    lineWrapping: true,
    mode: "javascript"
});
editor.setSize("100%", height);
editor.setValue(`option = {
    xAxis: {
      type: 'category',
      data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        data: [150, 230, 224, 218, 135, 147, 260],
        type: 'line'
      }
    ]
};`);

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
    var dom = document.getElementById("outputType")
    var outputType = dom.value;

    axios.post("/?outputType=" + outputType, data).then(function(resp) {
        if (outputType == "png") {
            document.getElementById("svg").innerHTML = '<img src="data:image/png;base64,' + resp + '" />';
        } else {
            document.getElementById("svg").innerHTML = resp;
        }
    }).catch(function(err) {
        alert(err.message);
    });
}