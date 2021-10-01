import React, { useContext } from "react";
import { LoadingUserInfoBar } from "../components/LoadingUserInfoBar";
import { UserInfoBar } from "../components/UserInfoBar";
import { useVideoListLatest } from "../hooks/api/useVideoListLatest";
import { usePlaylistList } from "../hooks/api/usePlaylistList";
import { useUser } from "../hooks/api/useUser";
import { useVideo } from "../hooks/api/useVideo";
import { Playlist, StatCard, Video } from "../shared/types";
import { AlertContext } from "../hooks/AlertContext";
import { useHistory } from "react-router";

export interface UserInfoBarControllerProps {
	className?: string;
}

export const UserInfoBarController: React.FC<UserInfoBarControllerProps> = ({ className }) => {
	const { pushAlert } = useContext(AlertContext);
	const history = useHistory();
	const [isLoadingUser, user] = useUser();
	const [isLoadingVideos, videos] = useVideoListLatest();
	const [isLoadingPlaylists, playlists] = usePlaylistList();
	const [, video] = useVideo(videos.length > 0 ? videos[0] : 0,  videos.length > 0);

	let isLoading = isLoadingUser || isLoadingPlaylists || isLoadingVideos;

	if (isLoading || user === null || playlists === null) {
		return <LoadingUserInfoBar></LoadingUserInfoBar>;
	}

	return (
		<UserInfoBar
			HomeClick={() => {
				pushAlert({
					message: "Coming soon...",
					type: "success",
				});
			}}
			LogOutClick={() => {
				history.push("/login");
			}}
			className={className}
			user={user}
			stats={ParseStats(playlists, video)}
		></UserInfoBar>
	);
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
			header: "Channels",
			count: channelCount,
			measurement: "total",
		},
		{
			header: "Last Updated",
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
