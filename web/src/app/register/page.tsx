"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function Register() {
    const router = useRouter();
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);
    async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        const form = new FormData(e.currentTarget);
        const firstName = form.get("firstName") as string;
        const lastName = form.get("lastName") as string;
        const username = form.get("username") as string;
        const email = form.get("email") as string;
        const password = form.get("password") as string;
        const dateOfBirth = form.get("dateOfBirth") as string;
        const about = form.get("about") as string;
        try{
        const res = await fetch("http://localhost:8080/api/auth/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ firstName, lastName, username, email, password, dateOfBirth, about }),
        });
        console.log(res);
        if (res.ok) {
            router.push("/login");
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
            <h1>Register</h1>
            <form onSubmit={handleSubmit}>
                <input name="firstName" type="text" placeholder="First Name" required/>
                <input name="lastName" type="text" placeholder="Last Name" required/>
                <input name="username" type="text" placeholder="Username (optional)" />
                <input name="email" type="email" placeholder="Email" required/>
                <input name="password" type="password" placeholder="Password" required/>
                <input name="dateOfBirth" type="date" placeholder="Date of Birth" required/>
                <input type="text" placeholder="About (optional)" />
                {error && <p>{error}</p>}
                <button type="submit" disabled={loading}>{loading ? "Registering..." : "Register"}</button>
                <Link href="/login"> Already have an account? Login</Link>
            </form>
        </div>
    );
}