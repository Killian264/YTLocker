import { useState } from "react";
import { usePlaylists } from "../hooks/api/usePlaylists";
import { PlaylistCreateCard } from "../components/PlaylistCreateCard";
import { PlaylistView } from "../components/PlaylistView";
import { useDeletePlaylist } from "../hooks/api/useDeletePlaylist";
import { useUpdatePlaylist } from "../hooks/api/useUpdatePlaylist";
import { usePlaylist } from "../hooks/api/usePlaylist";

export interface PlaylistViewControllerProps {
	className?: string;
	playlistId: number;
	BackClick: () => void;
}

export const PlaylistViewController: React.FC<PlaylistViewControllerProps> = ({
	className,
	playlistId,
	BackClick,
}) => {
	const [loadingPlaylistList, playlists] = usePlaylists();
	const [loadingPlaylist, playlist] = usePlaylist(playlistId);

	const [isEditing, setEditing] = useState(false);

	const deletePlaylist = useDeletePlaylist();
	const updatePlaylist = useUpdatePlaylist();

	if (loadingPlaylist || loadingPlaylistList || playlist == null) {
		return <div>Loading...</div>;
	}

	const swap = () => {
		setEditing(!isEditing);
	};

	if (isEditing) {
		return (
			<PlaylistCreateCard
				className={className}
				editPlaylist={playlist}
				playlists={playlists}
				CreateClick={(title, description, color) => {
					updatePlaylist(playlist.id, title, description, color);
				}}
				BackClick={swap}
			></PlaylistCreateCard>
		);
	}

	return (
		<PlaylistView
			className={className}
			playlist={playlist}
			EditClick={swap}
			DeleteClick={() => {
				deletePlaylist(playlistId);
				BackClick();
			}}
			BackClick={BackClick}
		></PlaylistView>
	);
};
