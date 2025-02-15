import { LoginForm } from "@/pages/auth/login/login-form";
import { Helmet } from "react-helmet-async";

export default function Login() {
  return (
    <>
      <Helmet title="login" />
      <LoginForm />
    </>
  );
}
