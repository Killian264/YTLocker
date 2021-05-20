import React from "react";
import { UserInfoBar } from "../components/UserInfoBar";
import { usePlaylists } from "../shared/api/usePlaylists";
import { useUser } from "../shared/api/useUser";
import { Playlist, StatCard } from "../shared/types";

export interface UserInfoBarControllerProps {
	className?: string;
}

export const UserInfoBarController: React.FC<UserInfoBarControllerProps> = ({ className }) => {
	const [userLoading, user] = useUser();
	const [playlistsLoading, playlists] = usePlaylists();

	let loading = userLoading || playlistsLoading;

	if (loading || user === null || playlists === null) {
		return <div>Loading...</div>;
	}

	return <UserInfoBar className="flex-grow" user={user} stats={ParseStats(playlists)}></UserInfoBar>;
};

const ParseStats = (playlists: Playlist[]): StatCard[] => {
	let videoCount = 0;
	let channelCount = 0;

	playlists.forEach((playlist) => {
		videoCount += playlist.videos.length;
		channelCount += playlist.channels.length;
	});

	return [
		{
			header: "Playlists",
			count: playlists.length,
			measurement: "total",
		},
		{
			header: "Videos",
			count: videoCount,
			measurement: "total",
		},
		{
			header: "Subscriptions",
			count: channelCount,
			measurement: "total",
		},
		{
			header: "Updated",
			count: 12,
			measurement: "hours ago",
		},
	];
};
