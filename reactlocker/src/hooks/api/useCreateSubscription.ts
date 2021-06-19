import axios from "axios";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { AlertContext } from "../AlertContext";

export const useCreateSubscription = (): ((playlistId: number, channelId: string) => Promise<void>) => {
	const { pushAlert } = useContext(AlertContext);
	const queryClient = useQueryClient();

	return async (playlistId: number, channelId: string) => {
		return axios
			.post(`/playlist/${playlistId}/subscribe/${channelId}`)
			.then(() => {
				pushAlert({
					message: "Subscription was created.",
					type: "success",
				});
				queryClient.invalidateQueries(["playlists"]);
				return;
			})
			.catch(() => {
				pushAlert({
					message: "Failed to create subscription",
					type: "failure",
				});
			});
	};
};
