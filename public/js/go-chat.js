$('#box-message').scrollTop($('#box-message')[0].scrollHeight);

$("#btn-photo").click(function () {
    $("#photo").click();
})
$("#btn-video").click(function () {
    $("#video").click()
})
// click video
$("#video").change(function () {
    let photo = $("#photo").val("");
    let match = ["video/mp4", "video/mov"];
    let file_data = $('#video').prop('files')[0];
    let ext = file_data.type;
    console.log(ext);
    if (ext != match[0] && ext != match[1]) {
        $("#video").val("");
        alert("Error: File isn't video!!!")
        return
    }
    let file = $("#video")[0].files;
    sentFileMedia(file);

})

// click photo
$("#photo").change(function () {
    let video = $("#video").val("");

    let match = ["image/gif", "image/png", "image/jpg","image/jpeg"];
    let file_data = $('#photo').prop('files')[0];
    let type = file_data.type;
    console.log(type);
    if (type != match[0] && type != match[1] && type != match[2] && type != match[3]) {
        $("#photo").val("");
        alert("Error: File isn't photo!!!")
        return
    }

    let file = $("#photo")[0].files;

    sentFileMedia(file);

})
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
            console.log(msg)
            location.reload();
        });

        request.fail(function (request, status, error) {
            console.log(request.responseText);
        });
    } else {
        alert("Please select a file.");
    }
}
// edit user name
$("#btn-edit-name").click(function () {
    let nameEdit = $("#name-edit").val(), idUser = $("#id").val();
    // Send the data using post
    let posting = $.post("/chat/edit-name", { id: idUser, name: nameEdit });

    // Put the results in a div
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
        msg = $("#message-input").val(),
        photo = $("#photo").val(),
        video = $("#video").val();
    let data = {};

    if (url == "/chat/edit-message") {
        let idMsg = $("#idMessage").val();
        console.log(idMsg);
        data = { idMessage: idMsg, newMessage: msg }
    } else {
        data = { id: userId, message: msg }
    }
    console.log(data);
    let posting = $.post(url, data);

    posting.done(function (mes) {
        console.log(mes)
        location.reload();
    });
    posting.fail(function (request, status, error) {
        console.log(request.responseText);
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
    url = "/chat/delete-message"
    let posting = $.post(url, { idMessage: id });

    posting.done(function (mes) {
        console.log(mes)
        location.reload();
    });
    posting.fail(function (request, status, error) {
        console.log(request.responseText);
    });

}