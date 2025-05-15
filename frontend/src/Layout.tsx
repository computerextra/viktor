import { NavLink, Outlet, useLocation } from "react-router";
import { AppSidebar } from "./components/app-sidebar";
import { Button } from "./components/ui/button";
import { SidebarInset, SidebarProvider } from "./components/ui/sidebar";

export default function Layout() {
  const location = useLocation();

  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <header className="sticky top-0 flex h-16 shrink-0 items-center gap-2 border-b bg-background px-4">
          <div className="border-b-1 w-full grid grid-cols-7 gap-0.5 p-1 items-center">
            <Button
              variant={location.pathname == "/" ? "default" : "outline"}
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
              <NavLink to="/">Mitarbeiter</NavLink>
            </Button>
            <Button
              variant={
                location.pathname.includes("/Lieferanten") ? "default" : "link"
              }
              asChild
            >
              <NavLink to="/">Lieferanten</NavLink>
            </Button>
            <Button
              variant={
                location.pathname.includes("/Archiv") ? "default" : "link"
              }
              asChild
            >
              <NavLink to="/">Archiv</NavLink>
            </Button>
            <Button
              variant={
                location.pathname.includes("/Suche") ? "default" : "link"
              }
              asChild
            >
              <NavLink to="/">Suche</NavLink>
            </Button>
            <Button
              variant={
                location.pathname.includes("/Werkstatt") ? "default" : "link"
              }
              asChild
            >
              <NavLink to="/">Werkstatt</NavLink>
            </Button>
          </div>
        </header>
        <div>
          <Outlet />
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
}
