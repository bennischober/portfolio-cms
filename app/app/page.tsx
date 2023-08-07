import { cookies } from "next/headers";
import { redirect } from "next/navigation";

async function checkAuth() {
  const cookie = cookies().get("jwt");
  return cookie;
}

export default async function Home() {
  return (
    <>
      Hello World!
    </>
  )
}
