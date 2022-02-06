<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" type="text/css" href="../asset/css/style.css">
</head>
<body>
  <h2>お問い合わせフォーム</h2>
  <form action="confirm" method="post">
    <table>
      <tr>
        <td>お名前</td>
        <td><input type="text" name="name"></td>
      </tr>
      <tr>
        <td>メールアドレス</td>
        <td><input type="email" name="email"></td>
      </tr>
      <tr>
        <td>お問い合わせ内容</td>
        <td><textarea name="inquiry" cols="40" rows="5" wrap="hard"></textarea></td>
      </tr>
    </table>
    <input type="submit" class="btn-push btn-push-blue" value="確認画面へ">
  </form>
</body>
</html>
