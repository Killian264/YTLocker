import { useState } from "react";
import { Footer } from "../components/Footer";
import { LinkButton } from "../components/LinkButton";
import { UserInfoBarController } from "../controllers/UserInfoBarController";
import { DashboardAccountsView } from "./DashboardAccountsView";
import { DashboardHelpView } from "./DashboardHelpView";
import { DashboardPlaylistView } from "./DashboardPlaylistsView";

export const DashboardPage: React.FC<{}> = () => {
	const [viewPage, setViewPage] = useState<"playlists" | "accounts" | "help">("playlists");

	let view = <DashboardAccountsView></DashboardAccountsView>;

	if (viewPage === "playlists") {
		view = <DashboardPlaylistView></DashboardPlaylistView>;
	}

	if (viewPage === "help") {
		view = <DashboardHelpView></DashboardHelpView>
	}

	return (
		<div className="h-screen flex flex-col flex-grow justify-between">
			<div className="mb-5 mt-6 pb-4">
				<div className="mb-4 px-4 mx-auto max-w-7xl">
					<UserInfoBarController></UserInfoBarController>
				</div>
				<div className="mb-4 px-4 mx-auto max-w-7xl">
					<div className="flex gap-6 mx-3">
						<LinkButton
							onClick={() => setViewPage("playlists")}
							selected={viewPage === "playlists"}
						>
							Playlists
						</LinkButton>
						<LinkButton
							onClick={() => setViewPage("accounts")}
							selected={viewPage === "accounts"}
						>
							Accounts
						</LinkButton>
						<LinkButton
							onClick={() => setViewPage("help")}
							selected={viewPage === "help"}
						>
							Help
						</LinkButton>
					</div>
				</div>
				<div className="px-4 mx-auto max-w-7xl">
					{view}
				</div>
			</div>
			<Footer></Footer>
		</div>
	);
};

