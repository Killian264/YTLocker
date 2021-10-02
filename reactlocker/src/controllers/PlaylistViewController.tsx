import { useState } from "react";
import { usePlaylistList } from "../hooks/api/usePlaylistList";
import { PlaylistCreateCard } from "../components/PlaylistCreateCard";
import { PlaylistView } from "../components/PlaylistView";
import { usePlaylistDelete } from "../hooks/api/usePlaylistDelete";
import { usePlaylistUpdate } from "../hooks/api/usePlaylistUpdate";
import { usePlaylist } from "../hooks/api/usePlaylist";
import { useAccountList } from "../hooks/api/useAccountList";
import { LoadingList } from "../components/LoadingList";
import { useAccount } from "../hooks/api/useAccount";

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
	const [isLoadingPlaylists, playlists] = usePlaylistList();
	const [isLoadingPlaylist, playlist] = usePlaylist(playlistId);
	const [isLoadingAccounts, accounts] = useAccountList();
	const [isLoadingAccount, account] = useAccount(playlist === null ? 0 : playlist.accountId);

	const [isEditing, setEditing] = useState(false);

	const deletePlaylist = usePlaylistDelete();
	const updatePlaylist = usePlaylistUpdate();

	if (
		isLoadingPlaylist ||
		isLoadingAccount ||
		isLoadingAccounts ||
		isLoadingPlaylists ||
		playlist == null ||
		account == null
	) {
		return <LoadingList limit={2}></LoadingList>;
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
				accounts={accounts}
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
			account={account}
			accounts={accounts}
			EditClick={swap}
			DeleteClick={() => {
				deletePlaylist(playlistId);
				BackClick();
			}}
			PauseClick={() => {}}
			CopyClick={() => {}}
			BackClick={BackClick}
		></PlaylistView>
	);
};
