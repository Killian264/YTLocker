import React from "react";
import { UserInfoBar } from "../components/UserInfoBar";
import { useLatestVideos } from "../shared/api/useLatestVideos";
import { usePlaylists } from "../shared/api/usePlaylists";
import { useUser } from "../shared/api/useUser";
import { useVideo } from "../shared/api/useVideo";
import { Playlist, StatCard, Video } from "../shared/types";

export interface UserInfoBarControllerProps {
	className?: string;
}

export const UserInfoBarController: React.FC<UserInfoBarControllerProps> = ({ className }) => {
	const [loadingU, user] = useUser();
	const [loadingP, playlists] = usePlaylists();
	const [loadingLV, videos] = useLatestVideos();
	const [loadingV, video] = useVideo(loadingLV ? 0 : videos[0]);

	let loading = loadingU || loadingP || loadingLV || loadingV;

	if (loading || user === null || playlists === null || videos === null || video === null) {
		return <div>Loading...</div>;
	}

	return <UserInfoBar className="flex-grow" user={user} stats={ParseStats(playlists, video)}></UserInfoBar>;
};

const ParseStats = (playlists: Playlist[], latestVideo: Video): StatCard[] => {
	let videoCount = 0;
	let channelCount = 0;

	var hours = Math.floor(Math.abs(new Date().getDate() - latestVideo.created.getDate()) / 36e5);

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
			count: hours,
			measurement: "hours ago",
		},
	];
};
