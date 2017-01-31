$('input:text, .ui.button', '.ui.action.input')
	.on('click', function(e) {
    	$('input:file', $(e.target).parents()).click();
	})
;

$('input:file', '.ui.action.input')
	.on('change', function(e) {
        var count = e.target.files.length;
    	var name = e.target.files[0].name + " + " + count +" other items";
    	$('input:text', $(e.target).parent()).val(name);
	});

function CheckFile(){
            var extension = ["mkv","flv","avi","mp4","mpg","mpeg","3gp"];
            var text ="";
            for(i=0;i<Object.keys($('#getfiles').get(0).files).length;i++){
                
                var fileName = $('#getfiles').get(0).files[i].name
                var fileExtension = fileName.substr(fileName.lastIndexOf(".")+1);
                if ($.inArray(fileExtension,extension) != -1){
                    text += fileName.substring(0,fileName.lastIndexOf(".")) + "$`&";
                }
            }
            document.getElementById('postnames').value = text;
}