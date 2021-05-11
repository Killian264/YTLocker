import React from "react";
import { Card } from "../components/Card";
import { ChannelListItem } from "../components/ChannelListItem";
import { PlaylistListItem } from "../components/PlaylistListItem";
import { PlusButton } from "../components/PlusButton";
import { UserInfoBar } from "../components/UserInfoBar";
import { VideoListItem } from "../components/VideosListItem";
import { Playlist, StatCard } from "../shared/types";

const user = {
	username: "Killian",
	email: "killiandebacker@gmail.com",
	joined: "Mar 13 2021",
};

const stats: StatCard[] = [
	{
		header: "Playlists",
		count: 454,
		measurement: "total",
	},
	{
		header: "Videos",
		count: 357,
		measurement: "total",
	},
	{
		header: "Subscriptions",
		count: 17,
		measurement: "total",
	},
	{
		header: "Updated",
		count: 13,
		measurement: "seconds ago",
	},
];

const playlists: Playlist[] = [
	{
		id: 932423423,
		youtube: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
		thumbnail:
			"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
		title: "DogeLog",
		description: "Videos showing Ben Awad as he builds dogehouse.",
		url:
			"https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK",
	},
];

export const DashboardPage: React.FC<{}> = () => {
	return (
		<>
			<div className="p-4 mx-auto max-w-7xl">
				<UserInfoBar
					className="flex-grow"
					user={user}
					stats={stats}
				></UserInfoBar>
			</div>
			<div className="px-4 mx-auto max-w-7xl flex">
				<Card className="w-7/12 mr-3">
					<div className="flex justify-between -mb-1 -mt-1">
						<div className="text-2xl font-semibold">
							<span className="leading-none -mt-0.5">
								Playlists
							</span>
						</div>
						<PlusButton
							color="secondary"
							disabled={false}
						></PlusButton>
					</div>
					<div>
						<PlaylistListItem
							playlist={playlists[0]}
						></PlaylistListItem>
						<PlaylistListItem
							playlist={playlists[0]}
							className="mt-2"
						></PlaylistListItem>
						<PlaylistListItem
							playlist={playlists[0]}
							className="mt-2"
						></PlaylistListItem>
					</div>
				</Card>
				<Card className="w-5/12">
					<div className="flex justify-between -mb-1 -mt-1">
						<div className="text-2xl font-semibold">
							<span className="leading-none -mt-0.5">Videos</span>
						</div>
					</div>
					<div>
						<VideoListItem video={playlists[0]}></VideoListItem>
						<VideoListItem
							video={playlists[0]}
							className="mt-2"
						></VideoListItem>
						<VideoListItem
							video={playlists[0]}
							className="mt-2"
						></VideoListItem>
						<VideoListItem
							video={playlists[0]}
							className="mt-2"
						></VideoListItem>
						<VideoListItem
							video={playlists[0]}
							className="mt-2"
						></VideoListItem>
					</div>
				</Card>
			</div>
			<div className="px-4 mx-auto max-w-7xl m-3">
				<Card>
					<div className="flex justify-between -mb-1 -mt-1">
						<div className="text-2xl font-semibold">
							<span className="leading-none -mt-0.5">
								Channels
							</span>
						</div>
					</div>
					<div>
						<ChannelListItem
							channel={playlists[0]}
						></ChannelListItem>
						<ChannelListItem
							channel={playlists[0]}
							className="mt-2"
						></ChannelListItem>
						<ChannelListItem
							channel={playlists[0]}
							className="mt-2"
						></ChannelListItem>
					</div>
				</Card>
			</div>
		</>
	);
};

DashboardPage.propTypes = {};
