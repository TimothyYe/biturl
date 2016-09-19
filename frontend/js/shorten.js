var Shorten = function() {
    var shortenMode = true;

    var validateUrl = function() {
        var reg = /(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?/;
        $('#txtUrl').on('input', function() {
            if (reg.test($(this).val())) {
                $('#submit-btn').prop('disabled', false);
            } else {
                $('#submit-btn').prop('disabled', true);
            }
        });
    };

    var handleShrink = function() {
        $('#submit-btn').on('click', function(e) {
            e.preventDefault();
            console.log("clicked!")

            if(shortenMode) {
              console.log("shrink mode");
              shortenMode = false;
              $('#submit-btn').text("COPY");
              //shrink_url();
            } else {
              console.log("copy mode");
              var clipboard = new Clipboard('.btn');

              clipboard.on('success', function(e) {
                console.log('Copied!');
                clipboard.destroy();
                shortenMode = true;
                $('#submit-btn').text("SHORTEN");
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
            validateUrl();
            handleShrink();
        }
    };

}();
