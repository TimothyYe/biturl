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

            if(shortenMode) {
              shrink_url();
            } else {
              var clipboard = new Clipboard('.btn');

              clipboard.on('success', function(e) {
                toastr.success('URL copied to clipboard!');
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
            if (data.result) {
              //Shorten success
                toastr.success('URL is shortened!');
                $('#txtUrl').prop('value', data.short);
                shortenMode = false;
                $('#submit-btn').text("COPY");
                $('#txtUrl').focus();
                $('#txtUrl').select();
            } else {
              //Shorten failed
                console.log('Shorten failed!');
                toastr.warning(data.message);
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
