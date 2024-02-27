$(document).ready(function(){
    $("#url").on('input', setBtnStatus);
    $("#code").on('input', setBtnStatus);
    $("#shorturl").on('input', setBtnStatus);

    $("#btn-save").on('click', () => {
        if (!$("#url").val().match("http(s)?://(.+)\.(.{2,})")) {
            Swal.fire(
                'Oops',
                'Invalid URL format',
                'error'
            );
            return;
        }

        const matchCode = $("#code").val().match(/[0-9]{6}/);
        if (!matchCode || matchCode[0] !== $("#code").val()) {
            Swal.fire(
                'Oops',
                'Invalid code format',
                'error'
            );
            return;
        }
        $("#btn-loading").css("display", "inline-grid");
        $("#btn-save").attr("disabled", true);
        let formData = new FormData();
        formData.append('short', $("#shorturl").val());
        formData.append('origin', btoa($("#url").val()));
        formData.append('code', $("#code").val());
        $.ajax({
            url : "/manual",
            type: "POST",
            data : formData,
            processData: false,
            contentType: false,
            success:function(data){
                $("#btn-loading").css("display", "none");
                $("#code").val("");
                $("#url").val("");
                $("#shorturl").val("");
                Swal.fire(
                    'Congratulations!',
                    data.message,
                    'success'
                )
            },
            error: function (err) {
                $("#btn-loading").css("display", "none");
                Swal.fire(
                    'Oops',
                    err.responseJSON.message,
                    'error'
                );
            }
        });
    });
})