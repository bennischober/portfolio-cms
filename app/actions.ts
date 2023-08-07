"use server";

import { cookies } from "next/headers";

export async function LogoutAction() {
    const req = await fetch("http://localhost:8080/api/auth/logout", {
        method: "POST",
    });

    console.log(req.json());

    if(req.ok) {
        cookies().delete("jwt");
    }
}
