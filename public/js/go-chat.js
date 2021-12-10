$(function () {
    // keep box message auto bottom
    $('#box-message').scrollTop($('#box-message')[0].scrollHeight);

    // load stamps list
    for (let index = 1; index <= 24; index++) {
        if (index < 10) {
            index = "0" + index;
        }
        let srcStamp = "/public/medias/stamps/" + index + ".png";
        $("#stamp-box").append("<button class=\"btn\" onclick=\"sentStamp('"+index+".png')\"><img src= " + srcStamp + "></button>")
    }
    // hide stamps list
    $("#stamp-box").hide()
    // show stamps list
    $("#btn-stamp").click(function () {
        $("#stamp-box").toggle()
    });
});

// send Stamp
const sentStamp = (id) => {
    let request = $.post('/chat/sent-stamp', {stamp: id})

    request.done(function () {
        location.reload();
    });

    request.fail(function (request, status, error) {
        alert(request.responseText);
    });

}

// redirect from button photo to input tag
$("#btn-photo").click(function () {
    $("#photo").click();
})

// redirect from button video to input tag
$("#btn-video").click(function () {
    $("#video").click()
})

// Upload video
$("#video").change(function () {
    $("#photo").val("");
    let match = ["video/mp4", "video/mov"];
    let file_data = $('#video').prop('files')[0];
    if (!match.includes(file_data.type)) {
        $("#video").val("");
        alert("Error: File isn't video!!!");
        return
    }
    let file = $("#video")[0].files;
    sentFileMedia(file);
})

// Upload photo
$("#photo").change(function () {
    $("#video").val("");

    let match = ["image/gif", "image/png", "image/jpg", "image/jpeg"];
    let file_data = $('#photo').prop('files')[0];

    if (!match.includes(file_data.type)) {
        $("#photo").val("");
        alert("Error: File isn't image!!!");
        return
    }

    let file = $("#photo")[0].files;
    sentFileMedia(file);

})

// ajax sent file
const sentFileMedia = (file) => {
    let fd = new FormData();

    if (file.length > 0) {
        fd.append('file', file[0]);

        let request = $.ajax({
            url: '/chat/upload-media',
            type: 'post',
            data: fd,
            contentType: false,
            processData: false,
        });
        request.done(function (msg) {
            console.log(msg);
            location.reload();
        });

        request.fail(function (request, status, error) {
            alert(request.responseText);
        });
    } else {
        alert("Please select a file.");
    }
}

// edit user name
$("#btn-edit-name").click(function () {
    let nameEdit = $("#name-edit").val(),
        idUser = $("#id").val();

    let posting = $.post("/chat/edit-name", { id: idUser, name: nameEdit });

    posting.done(function (mes) {
        alert(mes);
        $("#btn-close-moda-edit-name").click();
        $(".user-name").text(nameEdit);
    });
    posting.fail(function (request, status, error) {
        alert(request.responseText);
    });
})

// send message
$("#mesForm").submit(function (event) {
    event.preventDefault();
    let $form = $(this),
        url = $form.attr("action"),
        userId = $("#id").val(),
        msg = $("#message-input").val();
    let data = {};

    if (url == "/chat/edit-message") {
        let idMsg = $("#idMessage").val();
        data = { idMessage: idMsg, newMessage: msg };
    } else {
        data = { id: userId, message: msg };
    }

    let posting = $.post(url, data);

    posting.done(function (mes) {
        console.log(mes);
        location.reload();
    });
    posting.fail(function (request, status, error) {
        alert(request.responseText);
    });
})

// edit message
const editMessage = (id) => {
    let message = $("#msg-" + id).text();
    $("#message-input").val(message);
    $("#mesForm").attr("action", "/chat/edit-message");
    $("#idMessage").val(id);
    $("#cancel-edit-message").attr('hidden', false);;
}

// button cancel edit message
$("#cancel-edit-message").click(function () {
    $("#message-input").val("");
    $("#mesForm").attr("action", "/chat/sent-message");
    $("#idMessage").val("");
    $("#cancel-edit-message").attr('hidden', true);;
})

// delete message
const deleteMessage = (id) => {
    url = "/chat/delete-message";
    let posting = $.post(url, { idMessage: id });

    posting.done(function (mes) {
        console.log(mes);
        location.reload();
    });
    posting.fail(function (request, status, error) {
        alert(request.responseText);
    });

};