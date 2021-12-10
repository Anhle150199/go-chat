$("#loginForm").submit(function (event) {
    event.preventDefault();
    let $form = $(this),
        email = $form.find("input[name='email']").val(),
        password = $form.find("input[name='password']").val(),
        url = $form.attr("action");

    let posting = $.post(url, { email: email, password: password });
    posting.done(function (mes) {
        console.log(mes);
        window.location.href = '/chat/';
    });
    posting.fail(function (request, status, error) {
        alert(request.responseText);
    });
});

$("#form-register").submit(function (event) {
    event.preventDefault();
    let $form = $(this),
        name = $form.find("input[name='username']").val(),
        email = $form.find("input[name='email']").val(),
        password = $form.find("input[name='password']").val(),
        url = $form.attr("action");
    let data = { name: name, email: email, password: password };
    let posting = $.post(url, data);
    posting.done(function (mes) {
        alert(mes);
        window.location.href = '/user/login';
    });
    posting.fail(function (request, status, error) {
        alert(request.responseText);
    });

});