import { AppSidebar } from "@/components/app-sidebar";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { Logout } from "@api/userdata";
import useSession from "@hooks/useSession";
import { Reload } from "@wails/go/main/App";
import { NavLink, Outlet, useLocation } from "react-router";

export default function Layout() {
  const location = useLocation();
  const session = useSession();

  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <header className="sticky top-0 flex h-16 shrink-0 items-center gap-2 border-b bg-background px-4">
          <div className="w-full grid grid-cols-7 gap-0.5 items-center">
            <Button
              variant={location.pathname == "/" ? "default" : "link"}
              asChild
            >
              <NavLink to="/">Einkauf</NavLink>
            </Button>
            <Button
              variant={
                location.pathname.includes("/Mitarbeiter") ? "default" : "link"
              }
              asChild
            >
              <NavLink to="/Mitarbeiter">Mitarbeiter</NavLink>
            </Button>
            <Button
              variant={
                location.pathname.includes("/Lieferant") ? "default" : "link"
              }
              asChild
            >
              <NavLink to="/Lieferant">Lieferanten</NavLink>
            </Button>
            <Button
              variant={
                location.pathname.includes("/Archiv") ? "default" : "link"
              }
              asChild
            >
              <NavLink to="/Archiv">Archiv</NavLink>
            </Button>
            <Button
              variant={
                location.pathname.includes("/Suche") ? "default" : "link"
              }
              asChild
            >
              <NavLink to="/Suche">Suche</NavLink>
            </Button>
            <Button
              variant={
                location.pathname.includes("/Werkstatt") ? "default" : "link"
              }
              asChild
            >
              <NavLink to="/Werkstatt">Werkstatt</NavLink>
            </Button>
            {session?.Name ? (
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="link">{session?.Name ?? "Test"}</Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="w-56">
                  <DropdownMenuLabel>My Account</DropdownMenuLabel>
                  <DropdownMenuSeparator />
                  <DropdownMenuGroup>
                    <DropdownMenuItem>
                      {/* TODO: Link anpassen */}
                      <NavLink to="/">Profil</NavLink>
                      <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
                    </DropdownMenuItem>
                  </DropdownMenuGroup>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem
                    onClick={async () => {
                      const res = await Logout();
                      if (res) {
                        await Reload();
                      } else {
                        alert("Server Fehler!");
                      }
                    }}
                  >
                    Abmelden
                    <DropdownMenuShortcut>⇧⌘Q</DropdownMenuShortcut>
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            ) : (
              <Button variant={"link"} asChild>
                <NavLink to="/Anmelden">Anmelden</NavLink>
              </Button>
            )}
          </div>
        </header>
        <div>
          <Outlet />
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
}
