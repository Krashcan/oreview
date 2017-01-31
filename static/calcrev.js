$('input:text, .ui.button', '.ui.action.input')
	.on('click', function(e) {
    	$('input:file', $(e.target).parents()).click();
	})
;

$('input:file', '.ui.action.input')
	.on('change', function(e) {
        var count = e.target.files.length;
    	
	});

function CheckFile(){
            var extension = ["mkv","flv","avi","mp4","mpg","mpeg","3gp"];
            var text ="";
            var count = 0;
            for(i=0;i<Object.keys($('#getfiles').get(0).files).length;i++){
                
                var fileName = $('#getfiles').get(0).files[i].name
                var fileExtension = fileName.substr(fileName.lastIndexOf(".")+1);
                if ($.inArray(fileExtension,extension) != -1){
                    text += fileName.substring(0,fileName.lastIndexOf(".")) + "$`&";
                    count++;
                }
            }
            document.getElementById('movies').value = count + " videos";
            document.getElementById('postnames').value = text;
}