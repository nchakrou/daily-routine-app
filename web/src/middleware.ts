import { NextRequest, NextResponse } from "next/server";

export async function middleware(req: NextRequest) {
    const path = req.nextUrl.pathname;
    const token = req.cookies.get("token")?.value;
    if (!token) {
        if (path !== "/login" && path !== "/register") {
            return NextResponse.redirect(new URL("/login", req.url));
        }
        return NextResponse.next();
    }
    const res = await fetch("http://localhost:8080/api/auth/userinfo", {
        headers: {
            "Cookie": `token=${token}`
        }
    });
    if (!res.ok) {
        if (path !== "/login" && path !== "/register") {
            return NextResponse.redirect(new URL("/login", req.url));
        }
        return NextResponse.next();
    }
    if (path === "/login" || path === "/register") {
        return NextResponse.redirect(new URL("/", req.url));
    }
    return NextResponse.next();
}

export const config = {
    matcher: ["/"],
};