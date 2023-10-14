import { cookies } from "next/headers";
import { AuthenticationLogin, User } from "@/components/AuthenticationLogin";
import { redirect } from "next/navigation";

export const dynamic = "force-dynamic";

async function checkAuth() {
    const cookie = cookies().get("jwt");
    return cookie;
}

export default async function LoginPage() {
    const check = await checkAuth();
    if (check) redirect("/");

    const login = async (user: User) => {
        "use server";
        const response = await fetch(`${process.env.SERVER_HOST}/api/auth/login`, {
            method: "POST",
            body: JSON.stringify(user),
        });
        const data = await response.json();

        if (data.token) {
            cookies().set({
                name: "jwt",
                value: data.token,
                httpOnly: true,
                path: "/",
            });
        }

        return data;
    };

    return (
        <>
            <AuthenticationLogin login={login} />
        </>
    );
}
