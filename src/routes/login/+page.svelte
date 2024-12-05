<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "$lib/components/ui/card";
    import { goto } from "$app/navigation";

    let email = "";
    let password = "";

    async function handleLogin() {
        // TODO: Implement actual login logic here
        console.log("Login attempt with:", { email, password });
        
        // For now, just redirect to home page
        goto("/");
    }
</script>

<div class="min-h-screen flex items-center justify-center bg-background">
    <Card class="w-[350px]">
        <CardHeader>
            <CardTitle>Login</CardTitle>
            <CardDescription>Enter your credentials to access the application</CardDescription>
        </CardHeader>
        <CardContent>
            <form on:submit|preventDefault={handleLogin} class="space-y-4">
                <div class="space-y-2">
                    <label for="email" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
                        Email
                    </label>
                    <input
                        type="email"
                        id="email"
                        bind:value={email}
                        class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                        placeholder="m@example.com"
                        required
                    />
                </div>
                <div class="space-y-2">
                    <label for="password" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
                        Password
                    </label>
                    <input
                        type="password"
                        id="password"
                        bind:value={password}
                        class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                        required
                    />
                </div>
            </form>
        </CardContent>
        <CardFooter>
            <Button class="w-full" on:click={handleLogin}>
                Sign In
            </Button>
        </CardFooter>
    </Card>
</div>