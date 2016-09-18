var Shrink = function() {
    var shrinkMode = true;

    var handleShrink = function() {
        $('#submit-btn').on('click', function(e) {
            e.preventDefault();
            console.log("clicked!")

            if(shrinkMode) {
              console.log("shrink mode");
              shrinkMode = false;
              $('#submit-btn').text("Copy");
              //shrink_url();
            } else {
              console.log("copy mode");
              var clipboard = new Clipboard('.btn');

              clipboard.on('success', function(e) {
                console.log('Copied!');
                clipboard.destroy();
                shrinkMode = true;
                $('#submit-btn').text("Shrink URL");
              });

              clipboard.on('error', function(e) {
                console.log('Copy failed!');
              });
            }
        });
    }

    var shrink_url = function() {
        var form = $('#query-form');
        var url = form.attr("action");
        $.post(url, form.serialize(), function(data) {
            if (data.Code != '00') {
                bootbox.alert('');
            } else {
                bootbox.alert('');
            }
        }, "json");
    }

    return {
        //main function to initiate the module
        init: function() {
            handleShrink();
        }
    };

}();
