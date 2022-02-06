<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" type="text/css" href="../asset/css/style.css">
</head>
<body>
  <h2>ありがとうございます、以下の内容を送信しました</h2>
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
  <button type="button" class="btn-push" onclick="location.href='./form'">入力画面に戻る</button>
</body>
</html>
