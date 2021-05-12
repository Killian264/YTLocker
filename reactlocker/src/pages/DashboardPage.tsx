import React from "react";
import { RouteComponentProps } from "react-router-dom";
import { Card } from "../components/Card";
import { ChannelListItem } from "../components/ChannelListItem";
import { PlaylistListItem } from "../components/PlaylistListItem";
import { PlusButton } from "../components/PlusButton";
import { UserInfoBar } from "../components/UserInfoBar";
import { VideoListItem } from "../components/VideosListItem";
import { API } from "../shared/api";
import { useBearer } from "../shared/hooks/useBearer";
import { Channel, Playlist, StatCard, Video } from "../shared/types";

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
		created: new Date(),
	},
];

export const DashboardPage: React.FC<RouteComponentProps> = ({ history }) => {
	const [bearer, setBearer] = useBearer("");

	if (bearer == "") {
		history.push("/login");
	}

	// API.PlaylistList().then((res) => {
	// 	console.log(res);
	// });

	return (
		<>
			<div className="p-4 mx-auto max-w-7xl">
				<UserInfoBar
					className="flex-grow"
					user={user}
					stats={stats}
				></UserInfoBar>
			</div>
			<div className="px-4 mx-auto max-w-7xl grid grid-cols-12 gap-4">
				<PlaylistList
					className="xl:col-span-7 lg:col-span-6 col-span-12"
					playlists={[playlists[0], playlists[0]]}
				></PlaylistList>
				<VideoList
					className="xl:col-span-5 lg:col-span-6 col-span-12"
					videos={[
						playlists[0],
						playlists[0],
						playlists[0],
						playlists[0],
						playlists[0],
					]}
				></VideoList>
				<ChannelsList
					className="col-span-12"
					channels={[playlists[0], playlists[0], playlists[0]]}
				></ChannelsList>
			</div>
			<div className="px-4 mx-auto max-w-7xl m-3"></div>
		</>
	);
};

interface PlaylistListProps {
	className?: string;
	playlists: Playlist[];
}

export const PlaylistList: React.FC<PlaylistListProps> = ({
	className,
	playlists,
}) => {
	let list = playlists.map((playlist) => {
		return <PlaylistListItem playlist={playlist}></PlaylistListItem>;
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Playlists</span>
				</div>
				<PlusButton color="secondary" disabled={false}></PlusButton>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};

interface VideoListProps {
	className?: string;
	videos: Video[];
}

export const VideoList: React.FC<VideoListProps> = ({ className, videos }) => {
	let list = videos.map((video) => {
		return <VideoListItem video={video}></VideoListItem>;
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Videos</span>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};

interface ChannelsListProp {
	className?: string;
	channels: Channel[];
}

export const ChannelsList: React.FC<ChannelsListProp> = ({
	className,
	channels,
}) => {
	let list = channels.map((channel) => {
		return <ChannelListItem channel={channel}></ChannelListItem>;
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Channels</span>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};

DashboardPage.propTypes = {};
