$(document).ready(function(){
    const clipboard = new ClipboardJS('#btn-copy');
    clipboard.on('success', function () {
        Swal.fire('Copied!', 'URL has copied.', 'success');
    });
    window.sessionStorage.removeItem("res");
    let reload = false;

    $("#url").on('input', setBtnStatus);

    $("#btn-generate").on('click', () => {
        if (reload) {
            reload = false;
            window.location.reload();
            return;
        }
        if (!$("#url").val().match("http(s)?://(.+)\.(.{2,})")) {
            Swal.fire(
                'Oops',
                'Invalid URL format',
                'error'
            );
            return;
        }

        let formData = new FormData();
        const res = window.sessionStorage.getItem("res");
        formData.append('longUrl', btoa($("#url").val()));
        formData.append('token', res);
        window.sessionStorage.removeItem("res");
        $("#btn-loading").css("display", "inline-grid");
        $("#btn-generate").attr("disabled", true);
        $.ajax({
            url : "/short",
            type: "POST",
            data : formData,
            processData: false,
            contentType: false,
            success:function(data){
                reload = true;
                $("#btn-generate").html("Reload").attr("disabled", false);
                $("#btn-loading").css("display", "none");
                if (data.code === 200) {
                    $("#shorturl").val(data.url);
                    $("#btn-copy").attr("disabled", false);
                } else if (data.Code === 1) {
                    $("#shorturl").val(data.ShortUrl);
                    $("#btn-copy").attr("disabled", false);
                } else {
                    Swal.fire(
                        'Oops',
                        data.Message,
                        'error'
                    );
                }
            },
            error: function (err) {
                reload = true;
                $("#btn-generate").html("Reload").attr("disabled", false);
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