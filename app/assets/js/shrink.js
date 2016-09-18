var Shrink = function() {

    var handleShrink = function() {
        $('#submit-btn').on('click', function(e) {
            e.preventDefault();
            console.log("clicked!")
            //shrink_url();
        });
    }

    var shrink_url = function() {
        var form = $('#query-form');
        var url = form.attr("action");
        $.post(url, form.serialize(), function(data) {
            if (data.Code != '00') {
                $("#mobilePhone").val("");
                bootbox.alert('<b><font color="#d9534f">' + data.Message + '</font></b>');
            } else {
                $("#mobilePhone").val("");
                bootbox.alert('您查询的号码:<b><font color="#337ab7">' + data.Data.PhoneNumber + '</font></b> 为 <b><font color="#337ab7">' + data.Data.Comments + '</font></b>');
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
