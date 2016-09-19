var Shorten = function() {
    var shortenMode = true;

    var validateUrl = function() {
        var reg = /(https?:\/\/(?:www\.|(?!www))[^\s\.]+\.[^\s]{2,}|www\.[^\s]+\.[^\s]{2,})/;
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
        var form = $('#shorten-form');
        var url = form.attr("action");
        $.post(url, form.serialize(), function(data) {
            if (data.Result) {
              //Shorten success
                $('#txtUrl').prop('value', data.Short);
                return true;
            } else {
              //Shorten failed
                console.log('Shorten failed!');
              return false;
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
