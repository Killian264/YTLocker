import { Badge } from "../components/Badge";
import { Card } from "../components/Card";
import { GoogleOAuthCard } from "../components/GoogleOauthCard";
import { BeforeLinkingInfo, BaseAccountInfo } from "../components/InfoCards";
import { Logo } from "../components/Logo";
import { RightArrow } from "../components/Svg";
import { usePlaylistList } from "../hooks/api/usePlaylistList";

export const DashboardAccountsView: React.FC<{}> = () => {
	const [isLoadingPlaylists, playlists] = usePlaylistList();

	if (isLoadingPlaylists) {
		return <div>Loading...</div>;
	}

	const css = `hover:bg-primary-600 rounded-md flex justify-between cursor-pointer overflow-hidden`;
	const imageSize = "md:h-20 md:w-32 h-16 w-24";
	const textSize = "sm:text-md text-md";

	return (
		<div className="w-100">
			<div className="flex w-100 gap-8 flex-col lg:flex-row">
				<div className="flex-grow flex flex-col gap-8 lg:w-2/3">
					<Card>
						<div className="flex justify-between -mb-2 -mt-2">
							<div className="text-2xl font-semibold">
								<span className="leading-none -mt-0.5">Linked Accounts</span>
							</div>
						</div>
						<div>
							<div className={css} onClick={() => {}}>
								<div className="flex p-1 overflow-hidden">
									<div
										className={`rounded-lg flex-shrink-0 bg-black ${imageSize} flex justify-center items-center`}
									>
										<Logo></Logo>
									</div>
									<div className="pl-3 flex flex-col">
										<span className={`${textSize} font-semibold whitespace-nowrap`}>
											YTLocker Playlists |
											<span className="text-accent ml-2">{playlists.length}</span>/6
										</span>
										<span className="whitespace-nowrap whitespace-nowrap mt-0.5">
											<Badge className="mt-2 mr-2" color="primary">
												YTLocker
											</Badge>
											<Badge className="mr-2" color="secondary">
												MANAGE
											</Badge>
											<Badge color="secondary">
												LIMITED
											</Badge>
											<div className="mt-0.5">
												<span className="font-semibold">Email: </span> dev-locker@ytlocker.com
											</div>
										</span>
									</div>
								</div>
								<div className="mr-2 my-auto select-none">
									<RightArrow size={24}></RightArrow>
								</div>
							</div>
						</div>
					</Card>
					<BeforeLinkingInfo></BeforeLinkingInfo>
				</div>
				<div className="lg:w-1/3">
					<div className="flex flex-col md:flex-row lg:flex-col gap-8 items-start">
						<GoogleOAuthCard className="mx-auto flex-grow" type="link"></GoogleOAuthCard>
						<BaseAccountInfo></BaseAccountInfo>
					</div>
				</div>
			</div>
		</div>
	);
};