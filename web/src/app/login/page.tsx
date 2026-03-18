"use client";

import { useState } from "react";
import { useRouter } from 'next/navigation'
import Link from 'next/link'

export default function Login() {
    const router = useRouter();
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);
    async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        setLoading(true);
        const form = new FormData(e.currentTarget);
        const email = form.get("email") as string;
        const password = form.get("password") as string;
        try{
        const res = await fetch("http://localhost:8080/api/auth/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, password }),
        });
        if (res.ok) {
            router.push("/");
        } else {
            if (res.status == 500) {
                setError('Something went wrong, please try again later')
            } else {
                const data = await res.json();
                setError(data.error);
            }
        }
    }catch(e){
        setError('Something went wrong, please try again later')
    }finally{
        setLoading(false);
    }
    }
    return (
        <div>
            <h1>Login</h1>
            <form onSubmit={handleSubmit}>
                <input name="email" type="email" placeholder="Email" required onChange={() => setError("")}/>
                <input name="password" type="password" placeholder="Password" required onChange={() => setError("")}/>
                {error && <p>{error}</p>}
                <button type="submit" disabled={loading}>{loading ? "Logging in..." : "Login"}</button>
                <p>Don't have an account? <Link href="/register">Register</Link></p>
            </form>
        </div>
    );
}