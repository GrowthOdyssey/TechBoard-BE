package constants

// ステータス固定メッセージ
const UnauthorizedMessage = "認証エラー"
const NotFoundMessage = "そのURLは存在しません。"
const UnprocessableEntityErrorMessage = "項目毎のエラーが発生しています。"
const InternalServerErrorMessage = "サーバーエラー"

// 固有メッセージ
// ユーザー登録
const UserIdRequiredMessage = "ログインIDを入力してください。"
const UserIdMaxMessage = "ログインIDは20文字以内で入力してください。"
const PasswordRequiredMessage = "パスワードを入力してください。"
const PasswordMinMessage = "パスワードは8文字以上で入力してください。"
const PasswordMaxMessage = "パスワードは100文字以内で入力してください。"
const UserNameRequiredMessage = "ニックネームを入力してください。"
const UserNameMaxMessage = "ニックネームは100文字以内で入力してください。"
const ExistUser = "そのログインIDは既に使用されています。"
// ログイン
const LoginErrorMessage = "ログインIDかパスワードが異なります。"
