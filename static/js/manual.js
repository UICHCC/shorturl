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
                Swal.fire(
                    'Congratulations!',
                    data.message,
                    'success'
                )
            },
            error: function (err) {
                Swal.fire(
                    'Oops',
                    err.responseJSON.message,
                    'error'
                );
            }
        });
    });
})