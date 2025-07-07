import { Link } from "react-router";
import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuList,
  navigationMenuTriggerStyle,
} from "./ui/navigation-menu";
import { cn } from "@/lib/utils";

export default function NavBar() {
  return (
    <div className="w-full print:hidden">
      <NavigationMenu className="z-5 mx-auto">
        <NavigationMenuList>
          <NavigationMenuItem>
            <NavLink to="/" name="Start" />
          </NavigationMenuItem>
          <NavigationMenuItem>
            <NavLink to="/Einkauf" name="Einkauf" />
          </NavigationMenuItem>
          <NavigationMenuItem>
            <NavLink to="/Archiv" name="CE Archiv" />
          </NavigationMenuItem>
          <NavigationMenuItem>
            <NavLink to="/Warenlieferung" name="Warenlieferung" />
          </NavigationMenuItem>
          <NavigationMenuItem>
            <NavLink to="/CMS" name="CMS" />
          </NavigationMenuItem>
        </NavigationMenuList>
      </NavigationMenu>
    </div>
  );
}

function NavLink({ to, name }: { to: string; name: string }) {
  return (
    <Link
      to={to}
      className={cn(
        "hover:bg-sccent text-main-foreground rounded-base hover:border-border block space-y-1 border-2 border-transparent p-3 leading-none no-underline outline-hidden transition-colors select-none",
        navigationMenuTriggerStyle()
      )}
    >
      {name}
    </Link>
  );
}
