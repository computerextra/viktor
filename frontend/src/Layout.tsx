import { Outlet } from "react-router";
import NavBar from "./components/nav-bar";

export default function Layout() {
  return (
    <div>
      <NavBar />
      <div className="container mx-auto mt-5">
        <Outlet />
      </div>
    </div>
  );
}
