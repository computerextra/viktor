import { Link, Outlet } from "react-router";

export default function Layout() {
    return (
        <div>
            <Link to="/CMS">CMS</Link>
            <Outlet />
        </div>
    )
}