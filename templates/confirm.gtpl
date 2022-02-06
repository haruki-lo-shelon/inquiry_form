<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" type="text/css" href="../asset/css/style.css">
</head>
<body>
  <h2>お問い合わせ内容の確認</h2>
  <form action="send" method="post">
    <table>
      <tr>
        <td>お名前：</td>
        <td>{{.name}}</td>
      </tr>
      <tr>
        <td>メールアドレス：</td>
        <td>{{.email}}</td>
      </tr>
      <tr>
        <td>お問い合わせ内容：</td>
        <td>{{.inquiry}}</td>
      </tr>
    </table>
    <input type="hidden" name="name" value="{{.hid_name}}">
    <input type="hidden" name="email" value="{{.hid_email}}">
    <input type="hidden" name="inquiry" value="{{.hid_inquiry}}">
    <button type="button" class="btn-push" onclick="history.back()">入力画面に戻る</button>
    <input type="submit" class="btn-push btn-push-blue" value="この内容で登録する">
  </form>
</body>
</html>
