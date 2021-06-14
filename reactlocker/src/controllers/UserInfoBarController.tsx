import React from "react";
import { LoadingUserInfoBar } from "../components/LoadingUserInfoBar";
import { UserInfoBar } from "../components/UserInfoBar";
import { useLatestVideos } from "../hooks/api/useLatestVideos";
import { usePlaylists } from "../hooks/api/usePlaylists";
import { useUser } from "../hooks/api/useUser";
import { useVideo } from "../hooks/api/useVideo";
import { Playlist, StatCard, Video } from "../shared/types";

export interface UserInfoBarControllerProps {
	className?: string;
}

export const UserInfoBarController: React.FC<UserInfoBarControllerProps> = ({ className }) => {
	const [loadingU, user] = useUser();
	const [loadingP, playlists] = usePlaylists();
	const [loadingLV, videos] = useLatestVideos();
	const [, video] = useVideo(videos[0], !loadingLV);

	let loading = loadingU || loadingP || loadingLV;

	if (loading || user === null || playlists === null) {
		return <LoadingUserInfoBar></LoadingUserInfoBar>;
	}

	return <UserInfoBar className={className} user={user} stats={ParseStats(playlists, video)}></UserInfoBar>;
};

const ParseStats = (playlists: Playlist[], latestVideo: Video | null): StatCard[] => {
	let videoCount = 0;
	let channelCount = 0;

	playlists.forEach((playlist) => {
		videoCount += playlist.videos.length;
		channelCount += playlist.channels.length;
	});

	let [count, measurement] = [0, "hours ago"];

	if (latestVideo !== null) {
		[count, measurement] = UpdatedStats(latestVideo);
	}

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
			count: count,
			measurement: measurement,
		},
	];
};

const UpdatedStats = (video: Video): [number, string] => {
	const diff = Math.abs(new Date().valueOf() - video.created.valueOf());
	const hours = Math.floor(diff / 36e5);
	const minutes = Math.floor(diff / 6e4);
	const seconds = Math.floor(diff / 1e3);

	if (hours === 0 && minutes === 0) {
		return [seconds, "seconds ago"];
	}

	if (hours === 0) {
		return [minutes, "minutes ago"];
	}

	return [hours, "hours ago"];
};
